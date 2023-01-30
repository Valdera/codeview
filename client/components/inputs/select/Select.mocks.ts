import { ISelect } from './Select';

const base: ISelect = {
  data: [
    { value: 'easy_id', label: 'Easy', color: 'teal' },
    { value: 'medium_id', label: 'Medium', color: 'yellow' },
    {
      value: 'hard_id',
      label: 'Hard',
      color: 'red',
    },
  ],
  disabled: false,
  label: 'Select items',
};

export const mockSelectProps = {
  base,
};
