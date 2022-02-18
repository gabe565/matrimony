export const filterHouseholdQuestion = (isPartyHead, question) => {
  return isPartyHead || question.per !== "household";
};
