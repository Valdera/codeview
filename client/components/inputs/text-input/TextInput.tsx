import { Input } from '@mantine/core';
import { IconSearch } from '@tabler/icons';
import { ChangeEventHandler, useId } from 'react';

export interface ITextInput {
  value?: string;
  disabled?: boolean;
  label?: string;
  placeholder?: string;
  onChange?: ChangeEventHandler<HTMLInputElement>;
}

const TextInput: React.FC<ITextInput> = ({
  value,
  disabled = false,
  label = '',
  placeholder = '',
  onChange = () => {},
}) => {
  const id = useId();

  return (
    <Input.Wrapper id={id} label={label}>
      <Input
        id={id}
        value={value}
        onChange={onChange}
        disabled={disabled}
        placeholder={placeholder}
        icon={<IconSearch size={16} />}
        styles={(theme) => ({
          input: {
            backgroundColor: theme.colors.gray[1],
            color: theme.colors.gray[9],
          },
        })}
      />
    </Input.Wrapper>
  );
};

export default TextInput;
