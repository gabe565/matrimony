import { createStore } from "vuex";
import { filterHouseholdQuestion } from "../util/formatQuestion";

const { localStorage } = window;

const defaultPersistence = {
  query: {},
  party: {},
  responses: {},
  activeId: 0,
  currentQuestion: 0,
  responseSaved: {},
};

const saveState = (state) =>
  localStorage.setItem("state", JSON.stringify(state.persistent));
const loadState = () => JSON.parse(localStorage.getItem("state"));

export default createStore({
  state: () => ({
    persistent: loadState() || defaultPersistence,
    questions: [],
  }),
  mutations: {
    setQuery(state, query) {
      state.persistent.query = query;
      saveState(state);
    },
    setParty(state, party) {
      state.persistent.party = party;
      saveState(state);
    },
    setQuestions(state, questions) {
      state.questions = questions;
    },
    active(state, id) {
      state.persistent.activeId = id;
      if (!state.persistent.responses[id]) {
        state.persistent.responses[id] = {};
      }
      saveState(state);
    },
    setResponses(state, responses) {
      state.persistent.responses[state.persistent.activeId] = responses;
      saveState(state);
    },
    setSaved(state) {
      state.persistent.responseSaved[state.persistent.activeId] = true;
      state.persistent.activeId = 0;
      saveState(state);
    },
    addGuest(state, guest) {
      state.persistent.party.guests.push(guest);
      saveState(state);
    },
    resetPersistence(state) {
      state.persistent = defaultPersistence;
      saveState(state);
    },
  },
  getters: {
    activeGuest(state) {
      if (state.persistent.activeId) {
        return state.persistent.party.guests.find(
          (guest) => guest.id === state.persistent.activeId
        );
      }
      return {};
    },
    activeResponse(state, getters) {
      if (getters.activeGuest) {
        return state.persistent.responses[state.persistent.activeId];
      }
      return {};
    },
    visibleQuestions(state, getters) {
      let result = [];
      if (!state.persistent.activeId) {
        return result;
      }
      const { questions } = state;
      const isPartyHead =
        state.persistent.party.headId === state.persistent.activeId;
      const response = getters.activeResponse;
      for (const questionSection of questions) {
        const section = Object.entries(questionSection);
        result = result.concat(
          section.filter((e) => filterHouseholdQuestion(isPartyHead, e[1]))
        );

        for (const [type, question] of section) {
          const { field } = question;
          if (field in response && response[field]) {
            switch (type) {
              case "attendance": {
                const answer = question[response[field]];
                if (answer.followup) {
                  result = result.concat(
                    answer.followup
                      .map((q) => Object.entries(q))
                      .flat()
                      .filter((e) => filterHouseholdQuestion(isPartyHead, e[1]))
                  );
                }
                break;
              }
              case "choice": {
                const answer = question.choices.find(
                  (e) => e.answer === response[field]
                );
                if (answer && answer.followup) {
                  result = result.concat(
                    answer.followup
                      .map((q) => Object.entries(q))
                      .flat()
                      .filter((e) => filterHouseholdQuestion(isPartyHead, e[1]))
                  );
                }
                break;
              }
              default:
                break;
            }
          }
        }
      }
      return result;
    },
    allQuestionsAnswered(state, getters) {
      const questions = getters.visibleQuestions;
      const response = getters.activeResponse;
      return (
        response &&
        Object.values(response).filter((e) => !!e).length === questions.length
      );
    },
  },
  actions: {
    async fetchQuestions({ commit }) {
      const r = await fetch(`/api/rsvp/questions`);
      const questions = await r.json();
      commit("setQuestions", questions);
    },
    async fetchParty({ commit, state }) {
      const params = new URLSearchParams(
        Object.entries(state.persistent.query)
      );
      const r = await fetch(`/api/rsvp/init?${params}`);
      if (r.ok) {
        const party = await r.json();
        if (party.id !== state.persistent.party.id) {
          console.log("reset");
          commit("resetPersistence");
        }
        commit("setParty", party);
      }
      return r;
    },
    async respond(context, responses) {
      context.commit("setResponses", responses);
      const r = await fetch("/api/rsvp/response", {
        method: "PUT",
        body: JSON.stringify({
          values: responses,
          id: context.getters.activeGuest.id,
          sessionPassword: context.state.persistent.party.sessionPassword,
        }),
      });
      if (r.ok) {
        context.commit("setSaved");
      }
      return r;
    },
    async createGuest(context, guest) {
      const r = await fetch("/api/rsvp/guest/add", {
        method: "POST",
        body: JSON.stringify({
          guest,
          partyId: context.state.persistent.party.id,
          sessionPassword: context.state.persistent.party.sessionPassword,
        }),
      });
      if (r.ok) {
        const j = await r.json();
        context.commit("addGuest", j);
      }
      return r;
    },
  },
});
