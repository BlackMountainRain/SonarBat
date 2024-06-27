import React from 'react';
import TaskTable from '@/app/ui/tasks/table';

const TaskPage = (): React.ReactNode => (
  <div className="flex flex-col gap-5">
    <h1 className="text-xl font-semibold">All Tasks</h1>
    <TaskTable />
  </div>
);

export default TaskPage;
