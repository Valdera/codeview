import { Input, InputProps } from '@mantine/core';
import { ChangeEventHandler, ReactNode, useId } from 'react';

export interface ITextInput extends InputProps {
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
  ...props
}) => {
  const id = useId();

  return (
    <Input.Wrapper
      id={id}
      label={label}
      styles={(theme) => ({
        label: {
          marginBottom: '5px',
          color: theme.colors.gray[2],
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
        {...props}
      />
    </Input.Wrapper>
  );
};

export default TextInput;
