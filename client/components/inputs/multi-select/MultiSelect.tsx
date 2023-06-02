import { Flex, Text } from '@chakra-ui/react';
import {
  CloseButton,
  MultiSelect as MantineSelect,
  MultiSelectValueProps,
} from '@mantine/core';

/**TODO:
 * - Add Custom value
 */

export interface IMultiSelect {
  data: { value: string; label: string; color: string }[];
  value?: string[];
  creatable?: boolean;
  label?: string;
  placeholder?: string;
  disabled?: boolean;
  onChange?: (value: string[]) => void;
  onCreate?: ((query: string) => string | null) | undefined;
}

const MultiSelect: React.FC<IMultiSelect> = ({
  data,
  onCreate,
  onChange,
  value = [],
  disabled = false,
  creatable = true,
  label = '',
  placeholder = '',
}) => {
  return (
    <>
      <MantineSelect
        label={label}
        data={data}
        defaultValue={value}
        placeholder={placeholder}
        searchable
        creatable={creatable}
        valueComponent={(props) => {
          return <Value color={props.color} {...props} disabled={disabled} />;
        }}
        readOnly={disabled}
        transitionDuration={150}
        transition={'pop-top-left'}
        transitionTimingFunction={'ease'}
        getCreateLabel={(query) => `+ Create ${query}`}
        onCreate={onCreate}
        onChange={onChange}
        className={'w-full'}
        styles={(theme) => ({
          dropdown: {
            backgroundColor: theme.colors.background,
            borderColor: 'transparent',
            boxShadow: '1px 1px 25px -2px rgba(51,49,49,0.75)',
          },
          item: {
            color: `${theme.colors.gray[1]} !important`,
            '&:hover': {
              backgroundColor: '#374151',
            },
            '&[data-hovered]': {
              backgroundColor: '#374151',
            },
          },
          input: {
            backgroundColor: theme.colors.background,
            color: theme.colors.gray[1],
            width: '100%',
            fontWeight: 400,
            outline: 'none',
            borderColor: 'transparent',
            '&:focus': {
              borderColor: theme.colors.primary[7],
              transition: '.5s all',
            },
            '&:focus-within': {
              borderColor: theme.colors.primary[7],
              transition: '.5s all',
            },
          },
          wrapper: {
            width: '100%',
          },
          searchInput: {
            width: '100%',
          },
          label: {
            marginBottom: '5px',
            color: theme.colors.gray[2],
          },
        })}
      />
    </>
  );
};

const Value: React.FC<
  MultiSelectValueProps & { value: string; color: string; disabled?: boolean }
> = ({ color, label, onRemove, disabled, ...others }) => {
  return (
    <div className={others.className}>
      <Flex
        alignItems={'center'}
        justifyContent={'center'}
        backgroundColor={color}
        padding={'2px'}
        paddingLeft={'3'}
        paddingRight={disabled ? '3' : '0'}
        borderRadius={'md'}
      >
        <Text
          paddingY={disabled ? 1 : 0}
          color={'white'}
          lineHeight={'1'}
          fontWeight={'medium'}
        >
          {label}
        </Text>
        {!disabled && (
          <CloseButton
            onMouseDown={onRemove}
            variant={'transparent'}
            size={20}
            iconSize={16}
            tabIndex={-1}
            className={'text-white font-medium'}
          />
        )}
      </Flex>
    </div>
  );
};

export default MultiSelect;
