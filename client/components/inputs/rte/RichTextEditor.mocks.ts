import { IRichTextEditor } from './RichTextEditor';

const base: IRichTextEditor = {
  disabled: false,
  onSave: () => {},
};

export const mockRichTextEditorProps = {
  base,
};
