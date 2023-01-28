import { Button } from '@chakra-ui/react';
import { Select } from '@mantine/core';
import { Link, RichTextEditor as MantineRTE } from '@mantine/tiptap';
import BulletList from '@tiptap/extension-bullet-list';
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight';
import Highlight from '@tiptap/extension-highlight';
import ListItem from '@tiptap/extension-list-item';
import OrderedList from '@tiptap/extension-ordered-list';
import Paragraph from '@tiptap/extension-paragraph';
import SubScript from '@tiptap/extension-subscript';
import Superscript from '@tiptap/extension-superscript';
import TextAlign from '@tiptap/extension-text-align';
import Underline from '@tiptap/extension-underline';
import {
  NodeViewContent,
  NodeViewWrapper,
  ReactNodeViewRenderer,
  useEditor,
} from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';
import { lowlight } from 'lowlight';
import s from './RichTextEditor.module.scss';

// lowlight.registerLanguage('ts', tsSyntax);
// lowlight.registerLanguage('go', goSyntax);

// const content =
//   '<h2 style="text-align: center;">Welcome to Mantine rich text editor</h2><p><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a href="https://tiptap.dev/" rel="noopener noreferrer" target="_blank">Tiptap.dev</a> and supports all of its features:</p><ul><li>General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s> </li><li>Headings (h1-h6)</li><li>Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</li><li>Ordered and bullet lists</li><li>Text align&nbsp;</li><li>And all <a href="https://tiptap.dev/extensions" target="_blank" rel="noopener noreferrer">other extensions</a></li></ul>';

const content = `<h2 style="text-align: center">Welcome to Mantine rich text editor</h2><p><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a target="_blank" rel="noopener noreferrer nofollow" href="https://tiptap.dev/">Tiptap.dev</a> and supports all of its features:</p><ul><li><p>General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s></p></li><li><p>Headings (h1-h6)</p></li><li><p>Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</p></li><li><p>Ordered and bullet lists</p></li><li><p>Text align&nbsp;</p></li><li><p>And all <a target="_blank" rel="noopener noreferrer nofollow" href="https://tiptap.dev/extensions">other extensions</a></p></li></ul><pre><code>def positoin():
test</code></pre>`;

const RichTextEditor: React.FC = () => {
  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        orderedList: false,
        bulletList: false,
        paragraph: false,
      }),

      Paragraph.extend({
        addKeyboardShortcuts() {
          return {
            Tab: () => this.editor.commands.indent(),
            'Shift-Tab': () => this.editor.commands.outdent(),
          };
        },
      }),
      Underline,
      Link,
      Superscript,
      SubScript,
      Highlight,
      ListItem,
      OrderedList,
      BulletList,
      TextAlign.configure({ types: ['heading', 'paragraph'] }),
      CodeBlockLowlight.extend({
        addNodeView() {
          return ReactNodeViewRenderer(CodeSelection);
        },
        addKeyboardShortcuts() {
          return {
            // ↓ your new keyboard shortcut
            Tab: () => this.editor.commands.indent(),
            'Shift-Tab': () => this.editor.commands.outdent(),
          };
        },
      }).configure({ lowlight }),
    ],
    content,
  });
  return (
    <>
      <MantineRTE editor={editor}>
        <MantineRTE.Toolbar>
          <MantineRTE.ControlsGroup>
            <MantineRTE.Bold />
            <MantineRTE.Italic />
            <MantineRTE.Underline />
            <MantineRTE.Strikethrough />
            <MantineRTE.ClearFormatting />
            <MantineRTE.Highlight />
            <MantineRTE.Code />
          </MantineRTE.ControlsGroup>

          <MantineRTE.ControlsGroup>
            <MantineRTE.H1 />
            <MantineRTE.H2 />
            <MantineRTE.H3 />
            <MantineRTE.H4 />
          </MantineRTE.ControlsGroup>

          <MantineRTE.ControlsGroup>
            <MantineRTE.Blockquote />
            <MantineRTE.Hr />
            <MantineRTE.BulletList />
            <MantineRTE.OrderedList />
            <MantineRTE.Subscript />
            <MantineRTE.Superscript />
          </MantineRTE.ControlsGroup>

          <MantineRTE.ControlsGroup>
            <MantineRTE.Link />
            <MantineRTE.Unlink />
          </MantineRTE.ControlsGroup>

          <MantineRTE.ControlsGroup>
            <MantineRTE.AlignLeft />
            <MantineRTE.AlignCenter />
            <MantineRTE.AlignJustify />
            <MantineRTE.AlignRight />
          </MantineRTE.ControlsGroup>
        </MantineRTE.Toolbar>

        <MantineRTE.Content
          autoCorrect={'false'}
          spellCheck={'false'}
          className={'text-lg'}
        />
      </MantineRTE>
      <Button
        onClick={() => {
          console.log(editor?.getHTML());
        }}
      >
        Submit
      </Button>
    </>
  );
};

const CodeSelection = ({ updateAttributes, extension }: any) => (
  <NodeViewWrapper className={s['code-block']}>
    <Select
      size={'xs'}
      placeholder={'Language'}
      onChange={(value) => {
        updateAttributes({ language: value });
      }}
      defaultValue={'null'}
      className={'w-[10%] absolute right-3 top-3'}
      data={[
        { value: 'null', label: 'auto' },
        ...extension.options.lowlight
          .listLanguages()
          .map((lang: string) => ({ value: lang, label: lang })),
      ]}
    />

    <pre>
      <NodeViewContent as={'code'} />
    </pre>
  </NodeViewWrapper>
);

export default RichTextEditor;
