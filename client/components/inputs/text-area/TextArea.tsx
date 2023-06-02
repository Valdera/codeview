import { Textarea } from '@mantine/core';
import { ChangeEventHandler } from 'react';

export interface ITextInput {
  value?: string;
  disabled?: boolean;
  label?: string;
  placeholder?: string;
  onChange?: ChangeEventHandler<HTMLTextAreaElement>;
}

const TextArea: React.FC<ITextInput> = ({
  value,
  disabled = false,
  label = '',
  placeholder = '',
  onChange = () => {},
}) => {
  return (
    <Textarea
      value={value}
      onChange={onChange}
      disabled={disabled}
      placeholder={placeholder}
      label={label}
      styles={(theme) => ({
        label: {
          marginBottom: '5px',
          color: theme.colors.gray[2],
        },
        input: {
          backgroundColor: theme.colors.background,
          color: theme.colors.gray[2],
          fontWeight: 400,
          borderColor: 'transparent',
          outline: 'none',
          '&:focus': {
            borderColor: theme.colors.primary[7],
            transition: '.5s all',
          },
          '&:focus-within': {
            borderColor: theme.colors.primary[7],
            transition: '.5s all',
          },
        },
      })}
    />
  );
};

export default TextArea;
