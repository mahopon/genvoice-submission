import React, { useState } from 'react';
import { Table, message } from 'antd';
import { SurveyResponse } from '../types/Survey';
import SurveyDetailsModal from './SurveyDetailsModal';
import SurveyOwnerModal from './SurveyOwnerModal'; // Create this new component
import { useAuth } from '../context/AuthContext';

interface TableProps {
    surveys: SurveyResponse[];
    onSurveyComplete: () => void;
    onDeleteSurvey: (surveyId: string) => void;
}

const SurveyTable: React.FC<TableProps> = ({ surveys, onSurveyComplete, onDeleteSurvey }) => {
    const [messageApi, contextHolder] = message.useMessage();
    const [selectedSurvey, setSelectedSurvey] = useState<SurveyResponse | undefined>(undefined);
    const { userId } = useAuth();

    const handleSurveyClick = (record: SurveyResponse) => {
        setSelectedSurvey(record);
        console.log(record);
    };

    const columns = [
        {
            title: 'Survey Name',
            dataIndex: 'name',
            key: 'name',
            render: (text: string, record: SurveyResponse) => (
                <>
                    <a onClick={() => handleSurveyClick(record)} style={{ cursor: 'pointer' }}>
                        {text}
                    </a>
                </>
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
            dataIndex: 'created_by_name',
            key: 'created_by_name',
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

            {selectedSurvey && selectedSurvey.created_by === userId && (
                <SurveyOwnerModal
                    survey={selectedSurvey}
                    onClose={() => setSelectedSurvey(undefined)}
                    onDeleteSurvey={onDeleteSurvey}
                />
            )}
            {selectedSurvey && selectedSurvey.created_by !== userId && (
                <SurveyDetailsModal
                    selectedSurvey={selectedSurvey}
                    setSelectedSurvey={setSelectedSurvey}
                    onSurveyComplete={onSurveyComplete}
                    message={messageApi}
                />
            )}

        </>
    );
};

export default SurveyTable;
