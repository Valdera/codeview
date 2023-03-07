import { Input } from '@mantine/core';
import { ChangeEventHandler, ReactNode, useId } from 'react';

export interface ITextInput {
  value?: string;
  disabled?: boolean;
  label?: string;
  placeholder?: string;
  icon?: ReactNode;
  onChange?: ChangeEventHandler<HTMLInputElement>;
}

const TextInput: React.FC<ITextInput> = ({
  value,
  disabled = false,
  label = '',
  placeholder = '',
  icon = false,
  onChange = () => {},
}) => {
  const id = useId();

  return (
    <Input.Wrapper
      id={id}
      label={label}
      styles={(theme) => ({
        label: {
          marginBottom: '5px',
          color: theme.colors.primary[7],
        },
      })}
    >
      <Input
        id={id}
        value={value}
        onChange={onChange}
        disabled={disabled}
        placeholder={placeholder}
        icon={icon}
        styles={(theme) => ({
          input: {
            backgroundColor: '#EDF2F7',
            color: theme.colors.gray[9],
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
    </Input.Wrapper>
  );
};

export default TextInput;
