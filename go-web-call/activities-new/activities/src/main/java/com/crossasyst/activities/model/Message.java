package com.crossasyst.activities.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;
import java.sql.Timestamp;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Message {
    private String dataJobGuid;
    private Timestamp processingStatusType;
    private String logSessionId;
    private Timestamp processingStartDt;
    private Timestamp processingEndDt;
    private String attributes;
    private String subjectId;
    private String exceptionMessage;
    private String messageType;
    private String messageGuid;
    private String previousMessageGuid;
    private String externalPatientId;
    private Long portalConsumerId;
    private String externalProviderId;
    private Long portalStaffId;
    private String externalMessageId;
    private Integer revision;
    private String errorCd;
    private String errorDescription;
    private String errorSeverity;
}
