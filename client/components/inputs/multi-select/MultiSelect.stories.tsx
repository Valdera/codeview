import { ComponentMeta, ComponentStory } from '@storybook/react';
import { useState } from 'react';
import MultiSelect, { IMultiSelect } from './MultiSelect';
import { mockMultiSelectProps } from './MultiSelect.mock';

export default {
  title: 'inputs/MultiSelect',
  component: MultiSelect,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {},
} as ComponentMeta<typeof MultiSelect>;

// More on component templates: https://storybook.js.org/docs/react/writing-stories/introduction#using-args
const Template: ComponentStory<typeof MultiSelect> = (args) => {
  const [value, setValue] = useState<string[]>([]);

  return (
    <div className={'w-full m-10 flex items-center justify-center'}>
      <div style={{ width: '500px', padding: '10px' }}>
        <MultiSelect
          value={value}
          onChange={(val) => setValue(val)}
          {...args}
        />
      </div>
    </div>
  );
};

export const Base = Template.bind({});
// More on args: https://storybook.js.org/docs/react/writing-stories/args

Base.args = {
  ...mockMultiSelectProps.base,
} as IMultiSelect;
