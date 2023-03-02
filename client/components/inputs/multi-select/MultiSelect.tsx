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
          return <Value color={props.color} {...props} />;
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
          input: {
            backgroundColor: '#EDF2F7',
            color: theme.colors.gray[9],
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
            color: theme.colors.primary[7],
          },
        })}
      />
    </>
  );
};

const Value: React.FC<
  MultiSelectValueProps & { value: string; color: string }
> = ({ color, label, onRemove, ...others }) => {
  return (
    <div className={others.className}>
      <Flex
        alignItems={'center'}
        justifyContent={'center'}
        backgroundColor={color}
        padding={'2px'}
        paddingLeft={'3'}
        borderRadius={'md'}
      >
        <Text color={'white'} lineHeight={'1'} fontWeight={'medium'}>
          {label}
        </Text>
        <CloseButton
          onMouseDown={onRemove}
          variant={'transparent'}
          size={20}
          iconSize={12}
          tabIndex={-1}
          className={'text-white font-medium'}
        />
      </Flex>
    </div>
  );
};

export default MultiSelect;
