import React from 'react';
import { Button, Form, Spin } from '@douyinfe/semi-ui';
import { useWsPaginationTable } from '@hooks';
import { WsPaginationSearch, WsPaginationTable } from '@components';
import { editModal } from './components/edit';
import service from './service';
import { columns } from './columns';
import { IconPlus, IconSearch } from '@douyinfe/semi-icons';
import { useAsync } from 'react-use';

export default function ModelSettingsVisualEditor(props) {
  const table = useWsPaginationTable(service.getWsVideoRationPageList);
  const resolutionList = useAsync(service.getResolutionOptionsList);
  const resolutionOptions = resolutionList.value || [];

  return (
    <div>
      <Spin spinning={resolutionList.loading || table.tableProps.loading}>
        <WsPaginationSearch {...table.searchProps}>
          <Button
            icon={<IconPlus />}
            size='small'
            onClick={() => {
              editModal.open({}, () => {
                console.log('#');
              });
            }}
          >
            <>添加模型</>
          </Button>
          <Form.Input
            field='model_name'
            prefix={<IconSearch />}
            placeholder='搜索模型名称'
            showClear
            pure
            size='small'
          />
        </WsPaginationSearch>
        <WsPaginationTable {...table.tableProps} columns={columns(resolutionOptions)} />
      </Spin>
    </div>
  );
}
