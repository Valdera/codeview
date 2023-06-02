import { CheckIcon, EditIcon } from '@chakra-ui/icons';
import { IconButton } from '@chakra-ui/react';
import { Indent } from '@lib/tiptap/indent';
import { Select } from '@mantine/core';
import {
  Link,
  RichTextEditor as MantineRTE,
  useRichTextEditorContext,
} from '@mantine/tiptap';
import { IconBrandYoutube, IconPhotoUp } from '@tabler/icons';
import BulletList from '@tiptap/extension-bullet-list';
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight';
import Dropcursor from '@tiptap/extension-dropcursor';
import Highlight from '@tiptap/extension-highlight';
import Image from '@tiptap/extension-image';
import OrderedList from '@tiptap/extension-ordered-list';
import Paragraph from '@tiptap/extension-paragraph';
import SubScript from '@tiptap/extension-subscript';
import Superscript from '@tiptap/extension-superscript';
import TextAlign from '@tiptap/extension-text-align';
import Underline from '@tiptap/extension-underline';
import Youtube from '@tiptap/extension-youtube';
import {
  NodeViewContent,
  NodeViewWrapper,
  ReactNodeViewRenderer,
  useEditor,
} from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';
import axios from 'axios';
import cn from 'classnames';
import { lowlight } from 'lowlight';
import { ChangeEventHandler, ReactNode, useRef, useState } from 'react';
import s from './RichTextEditor.module.scss';

const defaultContent = `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Welcome to Code View</strong></h2><p style="margin-left: 0px!important;"><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a target="_blank" rel="noopener noreferrer nofollow" href="https://tiptap.dev/">Tiptap.dev</a> and supports all of its features</p><pre><code>def positoin():
test</code></pre>`;

/**TODO:
 * - Add animation
 * - Add documentation
 * - Move request to lib
 * - Add image resize (optional)
 */
export interface IRichTextEditor {
  content?: string;
  disabled?: boolean;
  onSave?: (content: string) => void;
  extraTools?: ReactNode[];
  heading?: ReactNode;
}

const RichTextEditor: React.FC<IRichTextEditor> = ({
  content = defaultContent,
  disabled = false,
  onSave = (_content) => {},
  extraTools = [],
  heading = undefined,
}) => {
  const [isLoading, setIsLoading] = useState(false);
  const [isEditable, setIsEditable] = useState(false);

  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        orderedList: false,
        bulletList: false,
        paragraph: false,
        codeBlock: false,
        dropcursor: false,
      }),
      Paragraph,
      Underline,
      Link,
      Superscript,
      Youtube,
      SubScript,
      Highlight,
      OrderedList,
      Dropcursor,
      BulletList,
      Image.configure({
        allowBase64: false,
      }),
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
    editable: !disabled && isEditable,
  });

  if (!editor) {
    return null;
  }

  const handleSubmit = () => {
    setIsLoading(true);

    onSave(editor.getHTML().toString());

    editor.setEditable(false);

    setIsEditable(false);
    setIsLoading(false);
  };

  const addImage = (files: FileList) => {
    if (files && files.length > 0) {
      for (const file of Array.from(files)) {
        const [mime] = file.type.split('/');

        if (mime === 'image') {
          let formData = new FormData();
          formData.append('imageFile', file);

          axios({
            method: 'post',
            url: 'http://localhost:8000/api/problem/image',
            data: formData,
            headers: { 'Content-Type': 'multipart/form-data' },
          })
            .then(function (response) {
              editor
                .chain()
                .focus()
                .setImage({ src: response.data.imageUrl })
                .run();
            })
            .catch(function (response) {
              console.log(response);
            });
        }
      }
    }
  };

  return (
    <>
      <MantineRTE
        editor={editor}
        className={cn(s.rte, 'relative')}
        styles={(theme) => ({
          root: {
            borderColor: 'transparent',
            backgroundColor: theme.colors.foreground,
          },
          toolbar: {
            paddingRight: '0',
            paddingLeft: '0',
            marginBottom: '5px',
            borderColor: 'transparent',
            backgroundColor: theme.colors.foreground,
          },
          toolbarControl: {
            backgroundColor: theme.colors.foreground,
          },
          controlsGroup: {
            color: 'white',
            backgroundColor: 'white',
            borderRadius: '5px',
          },
          control: {
            backgroundColor: 'white',

            // backgroundColor: theme.colors.background,
            // color: 'white',
          },
        })}
      >
        {!disabled && (
          <MantineRTE.Toolbar>
            {heading}

            {isEditable ? (
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
                  <InsertImageControl
                    handleFileChange={(evt) => {
                      if (evt.target.files) {
                        addImage(evt.target.files);
                      }
                    }}
                  />
                  <InsertYoutubeControl />
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
                  editor.setEditable(true);
                  setIsEditable(true);
                }}
              />
            )}
            {extraTools}
          </MantineRTE.Toolbar>
        )}

        <MantineRTE.Content
          autoCorrect={'false'}
          spellCheck={'false'}
          className={'text-lg'}
          onDrop={(e) => {
            e.preventDefault();
            e.stopPropagation();
            addImage(e.dataTransfer.files);
          }}
        />
      </MantineRTE>
    </>
  );
};

const InsertImageControl = ({
  handleFileChange,
}: {
  handleFileChange: ChangeEventHandler<HTMLInputElement>;
}) => {
  // const { editor } = useRichTextEditorContext();
  const inputRef: React.RefObject<HTMLInputElement> = useRef(null);

  return (
    <MantineRTE.Control
      onClick={() => inputRef.current?.click()}
      aria-label={'Insert image'}
      title={'Insert image'}
    >
      <input
        type={'file'}
        id={'file'}
        ref={inputRef}
        onChange={handleFileChange}
        style={{ display: 'none' }}
      />

      <IconPhotoUp stroke={1.5} size={16} />
    </MantineRTE.Control>
  );
};

const InsertYoutubeControl = () => {
  const { editor } = useRichTextEditorContext();

  const handleSubmit = () => {
    const url = prompt('Enter YouTube URL');
    if (url) {
      editor.commands.setYoutubeVideo({
        src: url,
        width: 640,
        height: 480,
      });
    }
  };

  return (
    <MantineRTE.Control
      onClick={handleSubmit}
      aria-label={'Insert image'}
      title={'Insert image'}
    >
      <IconBrandYoutube stroke={1.5} size={16} />
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
