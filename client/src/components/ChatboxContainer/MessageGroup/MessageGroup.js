import React from 'react';
import PropTypes from 'prop-types';
import Avatar from '../../Avatar/Avatar';
import Message from './Message/Message';

import "./MessageGroup.css";

function MessageGroup(props) {
    const { messages, right } = props;

    return (
        <div className={["MessageGroup", right ? "MessageGroup--Right" : "MessageGroup--Left"].join(" ")}>
            {!right && <div className="MessageGroup__Avatar">
                <Avatar />
            </div>}

            <div className={["MessageGroup__MessageContainer", right ? "MessageGroup__MessageContainer--Right" : "MessageGroup__MessageContainer--Left"].join(" ")}>
                {messages.map((message, index) => <Message content={message.content} createdAt={message.createdAt} key={index} />)}
            </div>
        </div>
    )
}

MessageGroup.propTypes = {
    right: PropTypes.bool
}

MessageGroup.defaultProps = {
    right: false
}

export default MessageGroup;