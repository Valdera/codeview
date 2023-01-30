import { ComponentMeta, ComponentStory } from '@storybook/react';
import { useState } from 'react';
import Select, { ISelect } from './Select';
import { mockSelectProps } from './Select.mocks';

export default {
  title: 'inputs/Select',
  component: Select,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {},
} as ComponentMeta<typeof Select>;

// More on component templates: https://storybook.js.org/docs/react/writing-stories/introduction#using-args
const Template: ComponentStory<typeof Select> = (args) => {
  const [value, setValue] = useState<string | null>('easy_id');

  return (
    <div className={'w-full m-10 flex items-center justify-center'}>
      <div style={{ width: '200px' }}>
        <Select value={value} handleChange={(val) => setValue(val)} {...args} />
      </div>
    </div>
  );
};

export const Base = Template.bind({});
// More on args: https://storybook.js.org/docs/react/writing-stories/args

Base.args = {
  ...mockSelectProps.base,
} as ISelect;
