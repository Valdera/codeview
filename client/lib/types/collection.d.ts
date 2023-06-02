import { Tag } from './metadata';

export type CollectionType = 'NOTE' | 'PROBLEM';

export interface Collection {
  id: string;
  title: string;
  tags: Tag[];
  emoji: string;
  type: CollectionType;
  description?: string;
  items?: CollectionItem[];
  createdAt?: string;
  updatedAt?: string;
}

export interface CollectionPreview {
  id: string;
  title: string;
  tags: Tag[];
  createdAt?: string;
  updatedAt?: string;
}

export interface CollectionItem {
  itemId: string;
  title: string;
  emoji: string;
  tags: Tag[];
  createdAt?: string;
  updatedAt?: string;
}
