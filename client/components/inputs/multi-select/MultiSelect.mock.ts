import { IMultiSelect } from './MultiSelect';

const base: IMultiSelect = {
  data: [
    { value: 'tag_1_id', label: 'Array', color: '#FC7300' },
    { value: 'tag_2_id', label: 'Binary Tree', color: '#00425A' },
    { value: 'tag_3_id', label: 'Graph', color: '#BFDB38' },
    { value: 'tag_4_id', label: 'Djikstra', color: '#F55050' },
    {
      value: 'tag_5_id',
      label: 'Topological Sort',
      color: '#FF78F0',
    },
    {
      value: 'tag_6_id',
      label: 'Binary Search',
      color: '#243763',
    },
    {
      value: 'tag_7_id',
      label: 'Two Pointers',
      color: '#61876E',
    },
  ],
  disabled: false,
  label: 'Select items',
};

export const mockMultiSelectProps = {
  base,
};
