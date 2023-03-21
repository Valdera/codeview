import { Source, Tag } from './metadata';

export interface Solution {
  id: string;
  content: string;
  createdAt?: string;
}

export interface Difficulty {
  id: string;
  label: string;
  color: string;
  createdAt?: string;
}

export interface Question {
  id: string;
  content: string;
  createdAt?: string;
}

export type Rating = 0 | 1 | 2 | 3 | 4 | 5;

export interface ProblemMetadata {
  title: string;
  difficulty: Difficulty;
  rating: Rating;
  sources: Source[];
  tags: Tag[];
}

export type Problem = ProblemMetadata & {
  id: string;
  slug?: string;
  question?: Question;
  solutions?: Solution[];
  createdAt?: string;
};
