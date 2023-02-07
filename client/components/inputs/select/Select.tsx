import { Tag, Text } from '@chakra-ui/react';
import { Group, Select as MantineSelect } from '@mantine/core';
import { forwardRef } from 'react';

type SupportedColor = 'teal' | 'yellow' | 'red';

export interface ISelect {
  value?: string | null;
  data: { value: string; label: string; color: SupportedColor }[];
  disabled?: boolean;
  label?: string;
  placeholder?: string;
  onChange?: (value: string | null) => void;
}

const Select: React.FC<ISelect> = ({
  value,
  data,
  onChange,
  disabled = false,
  label = '',
  placeholder = '',
}) => {
  const getColorFromValue = (value: string | null): string => {
    if (value) {
      for (const d of data) {
        if (d.value == value) {
          return d.color;
        }
      }
    }
    return 'gray';
  };

  return (
    <MantineSelect
      label={label}
      data={data}
      placeholder={placeholder}
      value={value}
      transition={'pop-top-left'}
      transitionDuration={80}
      itemComponent={SelectItem}
      transitionTimingFunction={'ease'}
      readOnly={disabled}
      onChange={onChange}
      styles={(theme) => ({
        item: {
          '&:[data-selected]': {
            color: 'red',
          },
          color: '#000',
          dataSelected: {
            color: 'red',
          },
          '.data-selected': { color: 'red' },
          '&:active': {
            color: 'red',
          },
          '&:data-selected': {
            color: 'red',
          },
        },
        itemsWrapper: {
          padding: '0',
        },
        input: {
          backgroundColor: value
            ? theme.colors[getColorFromValue(value)][1]
            : theme.colors.gray[1],
          color: value
            ? theme.colors[getColorFromValue(value)][9]
            : theme.colors.gray[9],
          fontWeight: 500,
          '&:focus': {
            outline: 'none',
            borderColor: 'transparent',
          },
          '&:focus-within': {
            outline: 'none',
            borderColor: 'transparent',
          },
        },
      })}
      className={'w-full'}
    />
  );
};

interface ItemProps extends React.ComponentPropsWithoutRef<'div'> {
  label: string;
  color: string;
  'data-selected': boolean;
}

const SelectItem = forwardRef<HTMLDivElement, ItemProps>(
  ({ color, label, ...others }: ItemProps, ref) => {
    const styleItem = others['data-selected']
      ? { backgroundColor: '#F1F3F4' }
      : {};
    return (
      <div
        ref={ref}
        onMouseEnter={others.onMouseEnter}
        onMouseDown={others.onMouseDown}
        style={styleItem}
        className={others.className}
      >
        <Group noWrap>
          <Tag
            colorScheme={color}
            width={'full'}
            display={'flex'}
            justifyContent={'center'}
          >
            <Text>{label}</Text>
          </Tag>
        </Group>
      </div>
    );
  }
);
SelectItem.displayName = 'SelectItem';

export default Select;
