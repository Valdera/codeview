import { Metadata, Problem, Question, Solution } from '@lib/types/problem';
import { StateCreator } from 'zustand';

export interface ProblemSlice {
  problem: Problem;
  createSolution: () => void;
  deleteSolution: (id: string) => void;
  updateSolution: (updatedSolution: Solution) => void;
  updateQuestion: (updatedQuestion: Question) => void;
  updateMetadata: (updatedMetadata: Partial<Metadata>) => void;
}

export const createProblemSlice: StateCreator<ProblemSlice> = (set, get) => {
  const initialProblem: Problem = {
    title: 'New Problem',
    difficulty: { id: '' },
    rating: 0,
    tags: [],
    sources: [],
    question: {
      content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Welcome to Code View</strong></h2>`,
    },
    solutions: [],
  };

  return {
    problem: initialProblem,
    createSolution: () => {
      const problem = get().problem;

      // TODO: create solution through backend

      const solution: Solution = {
        id: Date.now().toString(),
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Solution</strong></h2>`,
      };

      problem.solutions.push(solution);

      set({ problem });
    },
    deleteSolution: (id: string) => {
      const problem = get().problem;

      problem.solutions = problem.solutions.filter(
        (solution) => solution.id != id
      );

      set({ problem });
    },
    updateSolution: (updatedSolution: Solution) => {
      const problem = get().problem;

      problem.solutions = problem.solutions.map((solution) => {
        if (solution.id == updatedSolution.id) {
          return updatedSolution;
        }
        return solution;
      });

      set({ problem });
    },
    updateQuestion: (updatedQuestion: Question) => {
      const problem = get().problem;

      problem.question = updatedQuestion;

      set({ problem });
    },
    updateMetadata: (updatedMetadata: Partial<Metadata>) => {
      const problem = get().problem;

      console.log(updatedMetadata);

      set({ problem });
    },
  };
};
