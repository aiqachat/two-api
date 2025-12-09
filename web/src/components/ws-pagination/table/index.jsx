import { Card, Table } from '@douyinfe/semi-ui';
import { useEffectOnce } from 'react-use';

export const WsPaginationTable = ({
  pagination,
  tableData,
  reload,
  refresh,
  ...props
}) => {
  let showSizeChanger = true;
  if (pagination !== false && pagination?.showSizeChanger !== undefined) {
    showSizeChanger = pagination?.showSizeChanger;
  }
  const _pagination = {
    showQuickJumper: (pagination || undefined)?.showQuickJumper !== false,
    current: tableData.pageNumber,
    total: tableData.total,
    pageSize: tableData.pageSize,
    defaultPageSize: tableData.defaultPageSize,
    defaultCurrent: tableData.defaultCurrent,
    // showTotal: (total, range) =>
    //   `当前显示${range[0]}到${range[1]}, 共${total}条数据`,
    ...pagination,
    showSizeChanger,
  };

  useEffectOnce(() => {
    reload();
  });

  return (
    <Card>
      <Table
        {...props}
        dataSource={tableData.list}
        pagination={pagination === false ? false : _pagination}
      />
    </Card>
  );
};
