import { CheckIcon, EditIcon } from '@chakra-ui/icons';
import { Button, IconButton } from '@chakra-ui/react';
import { Indent } from '@lib/tiptap/indent';
import { Select } from '@mantine/core';
import {
  Link,
  RichTextEditor as MantineRTE,
  useRichTextEditorContext,
} from '@mantine/tiptap';
import { IconPhotoUp } from '@tabler/icons';
import BulletList from '@tiptap/extension-bullet-list';
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight';
import Dropcursor from '@tiptap/extension-dropcursor';
import Highlight from '@tiptap/extension-highlight';
import Image from '@tiptap/extension-image';
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
import { useCallback, useState } from 'react';
import s from './RichTextEditor.module.scss';

const content = `<h2 style="text-align: center; margin-left: 0px!important;"></h2><h2 style="text-align: center; margin-left: 0px!important;">Welcome to Mantine rich text editor</h2><p style="margin-left: 0px!important;"><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a target="_blank" rel="noopener noreferrer nofollow" href="https://tiptap.dev/">Tiptap.dev</a> and supports all of its features:</p><img src="https://www.shutterstock.com/image-photo/example-word-written-on-wooden-260nw-1765482248.jpg"><ul><li><p style="margin-left: 0px!important;">General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s></p></li><li><p style="margin-left: 0px!important;">Headings (h1-h6)</p></li><li><p style="margin-left: 0px!important;">Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</p></li><li><p style="margin-left: 0px!important;">Ordered and bullet lists</p></li><li><p style="margin-left: 0px!important;">Text align&nbsp;</p></li><li><p style="margin-left: 0px!important;">And all <a target="_blank" rel="noopener noreferrer nofollow" href="https://tiptap.dev/extensions">other extensions</a></p></li></ul><pre><code>def positoin():
test</code></pre>`;

export interface IRichTextEditor {
  disabled?: boolean;
  defaultEditable?: boolean;
  handleSave?: () => void;
}

const RichTextEditor: React.FC<IRichTextEditor> = ({
  defaultEditable = true,
  disabled = false,
}) => {
  const [isLoading, setIsLoading] = useState(false);
  const [editable, setEditable] = useState(defaultEditable);

  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        orderedList: false,
        bulletList: false,
        paragraph: false,
      }),
      Paragraph,
      Underline,
      Link,
      Superscript,
      SubScript,
      Highlight,
      ListItem,
      OrderedList,
      Image.configure({
        allowBase64: true,
      }),
      Dropcursor,
      BulletList,
      TextAlign.configure({ types: ['heading', 'paragraph'] }),
      CodeBlockLowlight.extend({
        addNodeView() {
          return ReactNodeViewRenderer(CodeSelection);
        },
        addKeyboardShortcuts() {
          return {
            Tab: () => this.editor.commands.insertContent('  '),
            'Shift-Enter': () => this.editor.commands.insertContent('\n'),
          };
        },
      }).configure({ lowlight }),
      Indent,
    ],
    content,
    editable: !disabled && editable,
  });

  const handleSubmit = () => {
    setIsLoading(true);

    editor?.setEditable(false);
    setEditable(false);
    setIsLoading(false);
  };

  const dropImage = (data: DataTransfer) => {
    const text = data.getData('text/plain');
    const { files } = data;
    console.log(files);

    if (files && files.length > 0) {
      for (const file of Array.from(files)) {
        const [mime] = file.type.split('/');

        if (mime === 'image') {
          const url = URL.createObjectURL(file);
          var reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onload = function () {
            console.log(reader.result);
            editor?.chain().focus().setImage({ src: reader.result }).run();
          };
          reader.onerror = function (error) {
            console.log('Error: ', error);
          };

          console.log('IMAGE URL  ' + url);
          // editor?.chain().focus().setImage({ src: url }).run();
        }
      }
    }
  };

  // Function must return a promise that resolves to uploaded image url.
  // After promise is resolved, blurred image placeholder with be replaced with the uploaded image's url.
  // Note that useCallback is required.
  const handleImageUpload = useCallback(
    (file: File): Promise<string> =>
      new Promise((_resolve, _reject) => {
        const formData = new FormData();
        formData.append('image', file);

        console.log(file);

        // fetch('https://api.imgbb.com/1/upload?key=api_key', {
        //   method: 'POST',
        //   body: formData,
        // })
        //   .then((response) => response.json())
        //   .then((result) => resolve(result.data.url))
        //   .catch(() => reject(new Error('Upload failed')));
        let url = 'https://source.unsplash.com/8xznAGy4HcY/800x400';
        if (url) {
          editor?.chain().focus().setImage({ src: url }).run();
        }
      }),
    [editor]
  );

  return (
    <>
      <MantineRTE editor={editor} className={'relative'}>
        {!disabled && (
          <MantineRTE.Toolbar>
            {editable ? (
              <>
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
                  <InsertImageControl />
                  <MantineRTE.Link />
                  <MantineRTE.Unlink />
                </MantineRTE.ControlsGroup>

                <MantineRTE.ControlsGroup>
                  <MantineRTE.AlignLeft />
                  <MantineRTE.AlignCenter />
                  <MantineRTE.AlignJustify />
                  <MantineRTE.AlignRight />
                </MantineRTE.ControlsGroup>

                <IconButton
                  className={'absolute ml-auto'}
                  aria-label={'Submit Text'}
                  icon={<CheckIcon />}
                  isLoading={isLoading}
                  onClick={handleSubmit}
                />
              </>
            ) : (
              <IconButton
                className={'absolute ml-auto'}
                aria-label={'Enable Edit Text'}
                isLoading={isLoading}
                icon={<EditIcon />}
                onClick={() => {
                  editor?.setEditable(true);
                  setEditable(true);
                }}
              />
            )}
          </MantineRTE.Toolbar>
        )}

        <MantineRTE.Content
          autoCorrect={'false'}
          spellCheck={'false'}
          className={'text-lg'}
          onDrop={(e) => {
            e.preventDefault();
            e.stopPropagation();
            console.log('DROP2');
            dropImage(e.dataTransfer);
          }}
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

const InsertImageControl = () => {
  const { editor } = useRichTextEditorContext();

  const addImage = useCallback(() => {
    const url = window.prompt('URL');

    if (url) {
      editor?.chain().focus().setImage({ src: url }).run();
    }
  }, [editor]);

  return (
    <MantineRTE.Control
      onClick={addImage}
      aria-label={'Insert image'}
      title={'Insert image'}
    >
      <IconPhotoUp stroke={1.5} size={16} />
    </MantineRTE.Control>
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
