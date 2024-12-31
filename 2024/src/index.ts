import One from "@/days/one";
import Two from "./days/two";
import Four from "./days/four";
import Five from "./days/five";
import { numberWords } from "./utils/days";

const day = parseInt(process.argv[2]);
const solutions = [One, Two, Four, Five];

switch (day) {
  case 1:
    One();
  case 2:
    Two();
  case 4:
    Four();
  case 5:
    Five();
  default:
    solutions[solutions.length - 1]();
}
