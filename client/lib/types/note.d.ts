import { Tag } from './metadata';

export type NoteStatus = 'DRAFT' | 'PUBLISHED';

export type Note = {
  id: string;
  title: string;
  tags: Tag[];
  emoji: string;
  status: NoteStatus;
  items?: NoteItem[];
  createdAt?: string;
  updatedAt?: string;
};

export interface NoteItem {
  id: string;
  numOrder: number;
  content: string;
  createdAt?: string;
  updatedAt?: string;
}
