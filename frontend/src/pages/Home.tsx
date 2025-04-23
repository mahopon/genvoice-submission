import React, { useEffect, useState } from 'react'
import { getSurveys } from '../api/SurveyAPI';
import SurveyTable from '../components/SurveyTable';
import { CollatedSurveyResponse, SurveyResponse } from '../types/Survey';
import AddSurveyButton from '../components/AddSurveyButton';
import { useAuth } from '../context/AuthContext';

const Home: React.FC = () => {
  const [userSurveys, setUserSurveys] = useState<SurveyResponse[]>([]);
  const [surveys, setSurveys] = useState<SurveyResponse[]>([]);
  const { isAuthenticated } = useAuth();

  const fetchSurveys = async () => {
    try {
      const fetchedSurveys: CollatedSurveyResponse = await getSurveys();
      setUserSurveys(fetchedSurveys.user_made);
      setSurveys(fetchedSurveys.others_made)
    } catch (error) {
      console.error("Error fetching surveys:", error);
    }
  };

  useEffect(() => {
    if (!isAuthenticated) return;
    fetchSurveys();
  }, [isAuthenticated]);

  return (
    <>
      {!isAuthenticated ?
        (<div style={{ maxWidth: "700px" }}>
          <h2>Writeup</h2>
          <p>Hi, I'm Chong Yao. I'm a year 2 at NTU reading Computer Science. I enjoy playing competitive FPS, Gacha games, and building keyboards (whenever time and money permits!). I like to think of myself as a sort of explorer, as I've explored different hobbies, games, or activities I think I might enjoy over the years, including 3D printing and soldering for fun. Recently, I've found that I enjoy web development, picking up Golang as my primary backend language, and recently started with React with Typescript for the frontend, with PostgreSQL as my database of choice and a touch of Redis as a cache-aside data store for faster retrievals. I've also touched a little bit on Python and Java for my school modules, having created a Flask backend with SQLite for a group project, and a CLI Java app while trying to adhere to SOLID principles as much as I can, implementing patterns such as repository, publisher-subscriber, and data transfer objects.

            Although my pool of projects might be small right now, I'm always looking to expand my knowledge and skills, even if it's a little bit everyday, with a growing list of things I want to work on and learn, such as gRPC, writing test code, and more about software design, I aim to become a well-rounded engineer in the future.
          </p>
        </div>)
        :
        (
          <>
            <div style={{ display: 'flex', justifyContent: 'flex-end', marginBottom: 16 }}>
              <AddSurveyButton refresh={fetchSurveys} />
            </div>
            <h4>User Created Surveys</h4>
            <SurveyTable surveys={userSurveys} onSurveyComplete={fetchSurveys} onDeleteSurvey={fetchSurveys} />
            <h4>Other Surveys</h4>
            <SurveyTable surveys={surveys} onSurveyComplete={fetchSurveys} onDeleteSurvey={fetchSurveys} />
          </>
        )
      }
    </>
  );
}

export default Home;
