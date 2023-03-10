import { Tag } from './metadata';

export interface NoteMetadata {
  title: string;
  tags: Tag[];
}

export type Note = NoteMetadata & {
  id: string;
  references: string[];
  items?: NoteItem[];
  slug?: string;
};

export interface NoteItem {
  id: string;
  numOrder: number;
  title: string;
  content: string;
}
