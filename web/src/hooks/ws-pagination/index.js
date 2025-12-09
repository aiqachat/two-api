import { useRef, useState } from 'react';
import { defaultTableData } from './const';
import { formatValues } from './utils';

export const useWsPaginationTable = (fetchFn, options) => {
  const _defaultPage = {
    pageNumber: 1,
    pageSize: 10,
    // 排序字段
    sortField: null,
    // 排序方式
    sortOrder: null,
    ...options,
  };
  const _options = options || {};
  if (_options.saveSearchParams !== true) {
    _options.saveSearchParams = false;
  }
  // 初始分页信息
  const initPageRef = useRef({
    ..._defaultPage,
  });

  /* 保存表格数据 */
  const [tableData, setTableData] = useState({
    ...defaultTableData,
    defaultPageSize: _defaultPage.pageSize,
    defaultCurrent: _defaultPage.pageNumber,
  });
  const [loading, setLoading] = useState(false);
  const loadTableData = async (params, showLoading = true) => {
    if (showLoading) {
      setLoading(true);
    }
    const res = await fetchFn({ ...params });
    if (showLoading) {
      setLoading(false);
    }
    if (!res) {
      return;
    }
    setTableData({ ...tableData, ...res });
  };

  const reload = (values, options) => {
    const showLoading = options?.showLoading !== false;
    initPageRef.current = { ..._defaultPage, ...values };
    loadTableData(initPageRef.current, showLoading).then();
  };

  // 表格数据变更监听
  const onChange = ({ pagination, sorter }) => {
    Object.assign(initPageRef.current, {
      pageNumber: pagination.current || _defaultPage.pageNumber,
      pageSize: pagination.pageSize || _defaultPage.pageSize,
    });
    const _sorter = sorter;
    if (_sorter?.order) {
      Object.assign(initPageRef.current, {
        sortField: _sorter.field,
        sortOrder: _sorter.order === 'descend' ? 'desc' : 'asc',
      });
    } else {
      Object.assign(initPageRef.current, {
        sortField: null,
        sortOrder: null,
      });
    }
    loadTableData(initPageRef.current).then();
  };

  const refresh = (options) => {
    const showLoading = options?.showLoading !== false;
    loadTableData(initPageRef.current, showLoading).then();
  };

  const onSearch = (values) => {
    const _values = formatValues(values);
    Object.assign(initPageRef.current, _defaultPage);
    Object.assign(initPageRef.current, _values);
    loadTableData(initPageRef.current).then();
  };

  return {
    tableProps: {
      tableData: tableData,
      loading,
      onChange,
      reload,
      refresh: () => refresh(),
    },
    searchProps: {
      onSearch,
      onReset: (values) => {
        reload(values);
      },
    },
    reload,
    refresh,
    updateList: (cb) => {
      setTableData({ ...tableData, list: cb(tableData.list) });
    },
    getSearchParams: () => {
      return {
        ...initPageRef.current,
      };
    },
  };
};
