import { Source, Tag } from './metadata';

export interface Solution {
  id: string;
  content: string;
  createdAt?: string;
  updatedAt?: string;
}

export interface Difficulty {
  id: string;
  label: string;
  color: string;
  createdAt?: string;
  updatedAt?: string;
}

export interface Question {
  id: string;
  content: string;
  createdAt?: string;
  updatedAt?: string;
}

export type Rating = 0 | 1 | 2 | 3 | 4 | 5;

export type Problem = {
  id: string;
  title: string;
  difficulty: Difficulty;
  rating: Rating;
  sources: Source[];
  tags: Tag[];
  question?: Question;
  solutions?: Solution[];
  createdAt?: string;
  updatedAt?: string;
};
