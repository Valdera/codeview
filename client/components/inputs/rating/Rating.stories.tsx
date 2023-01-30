import { ComponentMeta, ComponentStory } from '@storybook/react';
import { useState } from 'react';
import Rating, { IRating } from './Rating';
import { mockRatingProps } from './Rating.mocks';

export default {
  title: 'inputs/Rating',
  component: Rating,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {},
} as ComponentMeta<typeof Rating>;

// More on component templates: https://storybook.js.org/docs/react/writing-stories/introduction#using-args
const Template: ComponentStory<typeof Rating> = (args) => {
  const [value, setValue] = useState<number>(0);

  return (
    <div className={'w-full m-10 flex items-center justify-center'}>
      <Rating value={value} handleChange={(val) => setValue(val)} {...args} />
    </div>
  );
};

export const Base = Template.bind({});
// More on args: https://storybook.js.org/docs/react/writing-stories/args

Base.args = {
  ...mockRatingProps.base,
} as IRating;
