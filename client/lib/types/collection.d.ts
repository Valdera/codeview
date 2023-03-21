import { Tag } from './metadata';

export type CollectionType = 'NOTE' | 'PROBLEM';

export interface Collection {
  id: string;
  title: string;
  description?: string;
  tags: Tag[];
  imageUrl: string;
  collectionItems?: CollectionItem[];
  createdAt?: string;
}

export interface CollectionItem {
  type: CollectionType;
  targetId: string;
  slug?: string;
  createdAt?: string;
}
