import s from './ProblemListTablePagination.module.scss';

export interface IProblemListTablePagination {
  page: number;
  totalPage: number;
}

const ProblemListTablePagination: React.FC<IProblemListTablePagination> = ({
  page,
  totalPage,
}) => {
  return (
    <div className={'w-full sm:flex sm:items-center sm:justify-between '}>
      <div className={'text-sm text-gray-300 '}>
        Page{' '}
        <span className={'font-medium text-gray-100'}>
          {page} of {totalPage}
        </span>
      </div>

      <div className={'flex items-center mt-4 gap-x-4 sm:mt-0'}>
        <button className={s.button}>
          <svg
            xmlns={'http://www.w3.org/2000/svg'}
            fill={'none'}
            viewBox={'0 0 24 24'}
            strokeWidth={'1.5'}
            stroke={'currentColor'}
            className={'w-5 h-5 rtl:-scale-x-100'}
          >
            <path
              strokeLinecap={'round'}
              strokeLinejoin={'round'}
              d={'M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18'}
            />
          </svg>

          <span>previous</span>
        </button>

        <button className={s.button}>
          <span>Next</span>

          <svg
            xmlns={'http://www.w3.org/2000/svg'}
            fill={'none'}
            viewBox={'0 0 24 24'}
            strokeWidth={'1.5'}
            stroke={'currentColor'}
            className={'w-5 h-5 rtl:-scale-x-100'}
          >
            <path
              strokeLinecap={'round'}
              strokeLinejoin={'round'}
              d={'M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3'}
            />
          </svg>
        </button>
      </div>
    </div>
  );
};

export default ProblemListTablePagination;
