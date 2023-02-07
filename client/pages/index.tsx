import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';

import { NextPageWithLayout } from './page';

/**TODO:
 * - Add Affix (mantine)
 *
 */

const HomePage: NextPageWithLayout = () => {
  return <></>;
};

export default HomePage;

HomePage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};
