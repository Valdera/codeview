import { Flex, FlexProps, Icon, Link } from '@chakra-ui/react';
import React, { ReactNode } from 'react';
import { IconType } from 'react-icons';

interface INavigationSidebarLink extends FlexProps {
  icon: IconType;
  isActive?: boolean;
  children: ReactNode;
}

const NavigationSidebarLink: React.FC<INavigationSidebarLink> = ({
  icon,
  children,
  isActive = false,
  ...rest
}) => {
  return (
    <Link
      href={'#'}
      style={{ textDecoration: 'none', color: 'white' }}
      _focus={{ boxShadow: 'none' }}
    >
      <Flex
        align={'center'}
        fontWeight={'semibold'}
        fontSize={18}
        padding={4}
        marginX={4}
        marginY={3}
        borderRadius={'lg'}
        role={'group'}
        transition={'all .5s'}
        cursor={'pointer'}
        backgroundColor={isActive ? 'primary.500' : ''}
        _hover={{
          backgroundColor: 'primary.500',
        }}
        {...rest}
      >
        {icon && (
          <Icon
            marginRight={4}
            _groupHover={{
              color: 'white',
            }}
            as={icon}
          />
        )}
        {children}
      </Flex>
    </Link>
  );
};

export default NavigationSidebarLink;
