'use client';

import {
  getKeyValue,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from '@nextui-org/react';
import React from 'react';

const TaskTable = () => {
  const rows = [
    {
      key: '1',
      name: 'Tony Reichert',
      role: 'CEO',
      status: 'Active',
    },
    {
      key: '2',
      name: 'Zoey Lang',
      role: 'Technical Lead',
      status: 'Paused',
    },
    {
      key: '3',
      name: 'Jane Fisher',
      role: 'Senior Developer',
      status: 'Active',
    },
    {
      key: '4',
      name: 'William Howard',
      role: 'Community Manager',
      status: 'Vacation',
    },
  ];

  const columns = [
    {
      key: 'name',
      label: 'NAME',
    },
    {
      key: 'role',
      label: 'ROLE',
    },
    {
      key: 'status',
      label: 'STATUS',
    },
  ];

  return (
    <Table aria-label="Example table with dynamic content">
      <TableHeader columns={columns}>
        {(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
      </TableHeader>
      <TableBody items={rows}>
        {(item) => (
          <TableRow key={item.key}>
            {(columnKey) => (
              <TableCell>{getKeyValue(item, columnKey)}</TableCell>
            )}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
};

export default TaskTable;
