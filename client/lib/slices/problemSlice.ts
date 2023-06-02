import { Problem, Question, Solution } from '@lib/types';
import { StateCreator } from 'zustand';

export interface ProblemSlice {
  problem: Problem | null;
  load: (problem: Problem) => void;
  updateProblem: (body: {
    title: string;
    difficulty: string;
    rating: number;
    sources: string[];
    tags: string[];
  }) => void;
  createSolution: () => void;
  updateSolution: (id: string, body: { content: string }) => void;
  deleteSolution: (id: string) => void;
  createQuestion: () => void;
  updateQuestion: (body: { content: string }) => void;
}

export const createProblemSlice: StateCreator<ProblemSlice> = (set, get) => {
  return {
    problem: null,
    load: (problem) => {
      set({ problem });
    },
    updateProblem: (value) => {
      const problem = get().problem;
      if (!problem) return;

      // TODO: Update problem through backend

      // problem.title = title
      // problem.difficulty = difficulty;
      // problem.rating = rating;
      // problem.sources = sources;
      // problem.tags = tags;
      console.log('value to be post: ', value);

      set({ problem });
    },
    createSolution: () => {
      const problem = get().problem;
      if (!problem) return;

      // TODO: create solution through backend

      const solution: Solution = {
        id: Date.now().toString(),
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Solution</strong></h2>`,
      };

      problem.solutions?.push(solution);

      set({ problem });
    },
    deleteSolution: (id) => {
      const problem = get().problem;
      if (!problem || !problem.solutions) return;

      // TODO: delete solution through backend

      problem.solutions = problem.solutions.filter(
        (solution) => solution.id != id
      );

      set({ problem });
    },
    updateSolution: (id, { content }) => {
      const problem = get().problem;
      if (!problem || !problem.solutions) return;

      // TODO: Update solution through backend

      problem.solutions = problem.solutions.map((solution) => {
        if (solution.id == id) {
          return { ...solution, content: content };
        }
        return solution;
      });

      set({ problem });
    },
    createQuestion: () => {
      const problem = get().problem;
      if (!problem) return;

      // TODO: Create question through backend

      const question: Question = {
        id: Date.now().toString(),
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Solution</strong></h2>`,
      };

      problem.question = question;

      set({ problem });
    },
    updateQuestion: ({ content }) => {
      const problem = get().problem;
      if (!problem || !problem.question) return;

      problem.question.content = content;

      set({ problem });
    },
  };
};
