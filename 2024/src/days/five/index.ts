import getInput from "@/utils/get-input";

type ruleset = {
  base: number;
  before: number[];
  after: number[];
};

export default function Five() {
  console.log(import.meta.dirname);
  const data = getInput(import.meta.dirname);
  const rows = data.split("\n");

  let sectionFlag = false;

  let orderingRules = [];
  let updates = [];

  for (let row = 0; row < rows.length; row++) {
    if (rows[row].length < 2) {
      sectionFlag = true;
    }

    if (sectionFlag == false) {
      orderingRules.push(rows[row]);
    }

    if (sectionFlag === true && rows[row].length >= 2) {
      updates.push(rows[row]);
    }
  }

  // console.log(
  //   "Second section",
  //   updates.map((update) => update.split(","))
  // );

  const pairs = orderingRules.map((rule) => [
    parseInt(rule.split("|")[0]),
    parseInt(rule.split("|")[1]),
  ]);

  console.log("First section", pairs);

  const uniques = pairs
    .reduce((accumulator, current) => accumulator.concat(current), [])
    .filter((val, idx, arr) => arr.indexOf(val) === idx)
    .toSorted();

  const rulesets: Record<number, { before: number[]; after: number[] }> = {};
  uniques.map((num) => {
    rulesets[num] = { before: [], after: [] };
  });

  pairs.map((pair) => {
    rulesets[pair[0]].before.push(pair[1]);
    rulesets[pair[1]].after.push(pair[0]);
  });

  console.log(rulesets);
}
