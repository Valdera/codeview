import { ComponentMeta, ComponentStory } from '@storybook/react';
import RichTextEditor, { IRichTextEditor } from './RichTextEditor';
import { mockRichTextEditorProps } from './RichTextEditor.mocks';

export default {
  title: 'components/RichTextEditor',
  component: RichTextEditor,
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {},
} as ComponentMeta<typeof RichTextEditor>;

// More on component templates: https://storybook.js.org/docs/react/writing-stories/introduction#using-args
const Template: ComponentStory<typeof RichTextEditor> = (args) => (
  <RichTextEditor {...args} />
);

export const Base = Template.bind({});
// More on args: https://storybook.js.org/docs/react/writing-stories/args

Base.args = {
  ...mockRichTextEditorProps.base,
} as IRichTextEditor;
