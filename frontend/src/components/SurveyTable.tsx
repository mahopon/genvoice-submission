import React, { useState } from 'react';
import { Table, message } from 'antd';
import { SurveyResponse } from '../types/Survey';
import { useAuth } from '../context/AuthContext';
import SurveyDetailsModal from './SurveyDetailsModal';

interface TableProps {
    surveys: SurveyResponse[]
}

const SurveyTable: React.FC<TableProps> = ({ surveys }) => {
    const [messageApi, contextHolder] = message.useMessage();
    const { isAuthenticated } = useAuth();
    const [selectedSurvey, setSelectedSurvey] = useState<SurveyResponse | null>(null);

    const handleModalClose = () => {
        console.log("Close")
        // Do nothing
        setSelectedSurvey(null);
    }
    const handleModalOk = () => {
        console.log("OK")
        // TODO: Add POST to save answer
        setSelectedSurvey(null);
    }

    const columns = [
        {
            title: 'Survey Name',
            dataIndex: 'name',
            key: 'name',
            render: (text: string, record: SurveyResponse) => (
                <a onClick={() => {
                    if (!isAuthenticated)
                        messageApi.info("Please login to view and answer questions")
                    else
                        setSelectedSurvey(record);
                }} style={{ cursor: 'pointer' }}>
                    {text}
                </a>
            )
        },
        {
            title: 'Created Date',
            dataIndex: 'created_date',
            key: 'created_date',
            render: (text: string) => new Date(text).toLocaleDateString(),
        },
        {
            title: 'Created By',
            dataIndex: 'created_by',
            key: 'created_by',
        },
    ];

    return (
        <>
            {contextHolder}
            <Table
                dataSource={surveys}
                columns={columns}
                rowKey="id"
                pagination={false}
            />
            {selectedSurvey && (
                <SurveyDetailsModal selectedSurvey={selectedSurvey} onOk={handleModalOk} onClose={handleModalClose} />
            )}
        </>
    );
};

export default SurveyTable;
