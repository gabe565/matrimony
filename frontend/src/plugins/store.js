import { createStore } from "vuex";
import { filterHouseholdQuestion } from "../util/formatQuestion";
import { API_URL } from "../config/api";

const { localStorage } = window;

const defaultPersistence = {
  query: {},
  party: {},
  responses: {},
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
    activeGuestId: 0,
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
    setActiveGuestId(state, id) {
      if (id !== 0 && !state.persistent.responses[id]) {
        state.persistent.responses[id] = {};
      }
      state.activeGuestId = id;
      saveState(state);
    },
    setResponses(state, responses) {
      state.persistent.responses[state.activeGuestId] = responses;
      saveState(state);
    },
    setSaved(state) {
      state.persistent.responseSaved[state.activeGuestId] = true;
      saveState(state);
    },
    addGuest(state, guest) {
      state.persistent.party.guests.push(guest);
      saveState(state);
    },
    resetPersistence(state) {
      state.persistent = defaultPersistence;
      state.activeGuestId = 0;
      saveState(state);
    },
    replaceResponses(state, { responseSaved, responses }) {
      state.persistent.responseSaved = responseSaved;
      state.persistent.responses = responses;
      saveState(state);
    },
  },
  getters: {
    activeGuest(state) {
      if (state.activeGuestId) {
        return state.persistent.party.guests.find(
          (guest) => guest.id === state.activeGuestId
        );
      }
      return {};
    },
    partyHeadGuest(state) {
      if (state.persistent.party.guests.length) {
        return state.persistent.party.guests[0];
      }
      return {};
    },
    activeResponse(state, getters) {
      if (getters.activeGuestId) {
        return state.persistent.responses[state.persistent.activeGuestId];
      }
      return {};
    },
    visibleQuestions(state) {
      let result = [];
      if (!state.activeGuestId) {
        return result;
      }
      const { questions } = state;
      const isPartyHead = state.persistent.party.headId === state.activeGuestId;
      const response = state.persistent.responses[state.activeGuestId] || {};
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
      const r = await fetch(`${API_URL}/api/rsvp/questions`);
      const questions = await r.json();
      commit("setQuestions", questions);
    },
    async fetchParty({ commit, state }) {
      const params = new URLSearchParams(
        Object.entries(state.persistent.query)
      );
      const r = await fetch(`${API_URL}/api/rsvp/init?${params}`);
      const j = await r.json();
      if (r.ok) {
        if (j.id !== state.persistent.party.id) {
          commit("resetPersistence");
        }
        commit("setParty", j);
      }
      return j;
    },
    async fetchResponses(context) {
      const params = new URLSearchParams(
        Object.entries({
          id: context.state.persistent.party.id,
        })
      );
      const r = await fetch(`${API_URL}/api/rsvp/responses?${params}`, {
        headers: {
          Authorization: `Bearer ${context.state.persistent.party.sessionPassword}`,
        },
      });
      const j = await r.json();
      if (r.ok) {
        context.commit("replaceResponses", j);
        saveState(context.state);
      }
      return j;
    },
    async respond(context, responses) {
      context.commit("setResponses", responses);
      const r = await fetch(`${API_URL}/api/rsvp/response`, {
        method: "PUT",
        body: JSON.stringify({
          values: responses,
          id: context.state.activeGuestId,
          sessionPassword: context.state.persistent.party.sessionPassword,
        }),
      });
      if (r.ok) {
        context.commit("setSaved");
      }
      return r.json();
    },
    async createGuest(context, guest) {
      const r = await fetch(`${API_URL}/api/rsvp/guest/add`, {
        method: "POST",
        body: JSON.stringify({
          guest,
          partyId: context.state.persistent.party.id,
          sessionPassword: context.state.persistent.party.sessionPassword,
        }),
      });
      const j = await r.json();
      if (r.ok) {
        context.commit("addGuest", j);
      }
      return j;
    },
  },
});
