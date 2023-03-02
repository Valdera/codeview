export interface Source {
  id: string;
  label: string;
  color: string;
}

export interface Tag {
  id: string;
  label: string;
  color: string;
}

export interface Solution {
  id: string;
  content: string;
}

export interface Difficulty {
  id: string;
  label: string;
  color: string;
}

export interface Question {
  id: string;
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
  id: string;
  slug?: string;
  question?: Question;
  solutions?: Solution[];
};
