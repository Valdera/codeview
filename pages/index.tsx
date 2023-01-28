import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import RichTextEditor from '@components/rte/RichTextEditor';

import { NextPageWithLayout } from './page';

const Home: NextPageWithLayout = () => {
  return (
    <>
      <RichTextEditor />
    </>
  );
};

export default Home;

Home.getLayout = (page) => {
  return <PrimaryLayout>{page}</PrimaryLayout>;
};
