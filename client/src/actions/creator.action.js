import * as actionTypes from '../constants/actionTypes';

export const openModal = () =>
    ({ type: actionTypes.OPEN_CREATOR_MODAL });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_CREATOR_MODAL });