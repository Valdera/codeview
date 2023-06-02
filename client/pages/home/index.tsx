import { Icon } from '@chakra-ui/react';
import HomeFeatureListSection from '@components/home/sections/feature-list/HomeFeatureListSection';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import { NextPageWithLayout } from '@pages/page';
import { ReactElement } from 'react';
import {
  FcAbout,
  FcAssistant,
  FcCollaboration,
  FcDonate,
  FcManager,
} from 'react-icons/fc';

const HomePage: NextPageWithLayout = () => {
  return (
    <>
      <HomeFeatureListSection features={features} />
    </>
  );
};

HomePage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export default HomePage;

const features: {
  heading: string;
  description: string;
  icon: ReactElement;
  href: string;
}[] = [
  {
    heading: 'Problems',
    icon: <Icon as={FcAssistant} w={10} h={10} />,
    description: 'Solutions of programming interview questions',
    href: '/problem/list',
  },
  {
    heading: 'Notes',
    icon: <Icon as={FcCollaboration} w={10} h={10} />,
    description: 'Notes about software engineering (BE, FE, ML, etc)',
    href: '/note/list',
  },
  {
    heading: 'Tutorials',
    icon: <Icon as={FcDonate} w={10} h={10} />,
    description: 'Tutorial video about programming',
    href: '/tutorial/list',
  },
  {
    heading: 'Blogs',
    icon: <Icon as={FcManager} w={10} h={10} />,
    description: 'Published blogs on Medium',
    href: '/blog/list',
  },
  {
    heading: 'Collections',
    icon: <Icon as={FcAbout} w={10} h={10} />,
    description: 'Collections of multiple source',
    href: '/collection/list',
  },
];
