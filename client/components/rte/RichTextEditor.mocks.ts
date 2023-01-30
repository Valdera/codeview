import { IRichTextEditor } from './RichTextEditor';

const base: IRichTextEditor = {
  disabled: false,
  defaultEditable: true,
  handleSave: () => {},
};

export const mockRichTextEditorProps = {
  base,
};
