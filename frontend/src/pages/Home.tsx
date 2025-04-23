import React, { useEffect, useState } from 'react'
import { getSurveys } from '../api/SurveyAPI';
import SurveyTable from '../components/SurveyTable';
import { SurveyResponse } from '../types/Survey';

const Home: React.FC = () => {
  const [surveys, setSurveys] = useState<SurveyResponse[]>([]);

  const fetchSurveys = async () => {
    try {
      const fetchedSurveys = await getSurveys();
      setSurveys(fetchedSurveys);
    } catch (error) {
      console.error("Error fetching surveys:", error);
    }
  };

  useEffect(() => {
    fetchSurveys();
  }, []);

  return (
    <SurveyTable surveys={surveys} />
  );
}

export default Home