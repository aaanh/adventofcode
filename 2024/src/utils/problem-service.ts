import { exit } from "process";
import { numberWords } from "./days";
import fs from "fs";

async function getProblemInput(
  token: string,
  day: number,
  year: number,
  part: number = 1
) {
  if (fs.existsSync(`src/days/${numberWords[day - 1]}/input-${part}.txt`)) {
    return Promise.resolve("");
  }

  const url = `https://adventofcode.com/${year}/day/${day}/input`;
  const headers = {
    cookie: `session=${token}`,
  };

  return fetch(url, { headers })
    .then((res) => res.text())
    .then((data) => data);
}

function createProblemFolder(problemPath: string, problemIndex: string) {
  if (!fs.existsSync(problemPath)) {
    fs.mkdirSync(problemPath);
  }

  if (!fs.existsSync(problemIndex)) {
    fs.writeFileSync(problemIndex, "");
  }
}

async function ProblemService() {
  console.log(process.argv);

  const token = process.env.AOC_TOKEN;

  if (!token) {
    throw new Error("Advent of Code user token is not supplied.");
  }

  const today = new Date().getDate();
  const year = new Date().getFullYear();
  let day = numberWords[today - 1];

  if (process.argv[2]) {
    const argvDay = parseInt(process.argv[2]);
    day = numberWords[argvDay - 1];
    const problemPart = parseInt(process.argv[3] ?? 1);
    const problemPath = `src/days/${day}`;
    const problemIndex = `${problemPath}/index.ts`;
    const problemInput = `${problemPath}/input.txt`;
    createProblemFolder(problemPath, problemIndex);
    await getProblemInput(token, argvDay, year, problemPart).then((input) => {
      fs.writeFileSync(problemInput, input);
    });

    exit(0);
  }

  const problemPart = process.argv[3] ?? 1;
  const problemPath = `src/days/${day}`;
  const problemIndex = `${problemPath}/index.ts`;
  const problemInput = `${problemPath}/input.txt`;
  createProblemFolder(problemPath, problemIndex);
  await getProblemInput(token, today, year).then((input) => {
    fs.writeFileSync(problemInput, input);
  });

  exit(0);
}

await ProblemService();
