import { createStore } from "vuex";

const { localStorage } = window;

const defaultPersistence = {
  query: {},
  party: {},
  responses: {},
  activeId: 0,
  currentQuestion: 0,
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
      const response = getters.activeResponse;
      for (const questionSection of questions) {
        const section = Object.entries(questionSection);
        result = result.concat(section);
        for (const [type, question] of section) {
          const { field } = question;
          if (field in response && response[field]) {
            switch (type) {
              case "attendance": {
                const answer = question[response[field]];
                if (answer.followup) {
                  result = result.concat(
                    answer.followup.map((q) => Object.entries(q)).flat()
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
                    answer.followup.map((q) => Object.entries(q)).flat()
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
      console.log(Object.values(response).filter((e) => !!e).length);
      console.log(questions.length);
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
      const party = await r.json();
      commit("setParty", party);
      return r.ok;
    },
    async respond(context, responses) {
      context.commit("setResponses", responses);
      return await fetch("/api/rsvp/response", {
        method: "PUT",
        body: JSON.stringify({
          values: responses,
          id: context.getters.activeGuest.id,
          sessionPassword: context.state.persistent.party.sessionPassword,
        }),
      });
    },
  },
});
