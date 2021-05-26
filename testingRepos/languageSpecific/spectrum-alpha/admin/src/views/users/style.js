import styled from 'styled-components';
import { FlexCol } from '../../components/globals';

export const View = styled(FlexCol)`
  width: 100%;
  height: 100%;
`;

export const UserCommunitySettingsContainer = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1 0 auto;
  align-items: center;
`;

export const SectionHeader = styled.h1`
  font-size: 24px;
  font-weight: 800;
  padding: 32px 16px 16px;
  background: ${props => props.theme.bg.default};
  color: ${props => props.theme.text.default};
`;
