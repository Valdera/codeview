export interface Source {
  id: string;
}

export interface Tag {
  id: string;
}

export interface Solution {
  id: string;
  content: string;
}

export interface Difficulty {
  id: string;
}

export interface Question {
  content: string;
}

export type Rating = 0 | 1 | 2 | 3 | 4 | 5;

export interface Metadata {
  title: string;
  difficulty: Difficulty;
  rating: Rating;
  sources: Source[];
  tags: Tag[];
}

export type Problem = Metadata & {
  question: Question;
  solutions: Solution[];
};
