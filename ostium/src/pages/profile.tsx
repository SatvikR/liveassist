import React from "react";
import { AuthPage } from "../components/AuthPage";
import { PageLayout } from "../components/PageLayout";
import { Profile } from "../ui/users/Profile";

export interface ProfilePageProps {}

const ProfilePage: React.FC<ProfilePageProps> = ({}) => {
  return (
    <PageLayout withHead>
      <AuthPage>
        <Profile />
      </AuthPage>
    </PageLayout>
  );
};

export default ProfilePage;
