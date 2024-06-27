import React from 'react';
import ResourceTable from '@/app/ui/resources/table';

const ResourcePage = (): React.ReactNode => {
  return (
    <div className="flex flex-col gap-5">
      <h1 className="text-xl font-semibold">All Resources</h1>
      <ResourceTable />
    </div>
  );
};

export default ResourcePage;
