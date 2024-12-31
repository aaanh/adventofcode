import fs from "fs";

export default function getInput(dirname: string) {
  const data = fs.readFileSync(`${dirname}/input.txt`, "utf-8");

  return data;
}
