import { IRichTextEditor } from './RichTextEditor';

const base: IRichTextEditor = {
  disabled: false,
  defaultEditable: true,
  onSave: () => {},
};

export const mockRichTextEditorProps = {
  base,
};
