package com.crossasyst.activities.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.format.annotation.DateTimeFormat;

import javax.persistence.*;
import java.sql.Timestamp;
import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "message")
public class MessageEntity {

    @Id
    @Column(name = "msg_Id")
    private Long msgId;
    @Column(name = "data_job_guid")
    private String dataJobGuid;
    @Column(name = "processing_status_type")
    private Timestamp processingStatusType;
    @Column(name = "log_session_id")
    private String logSessionId;
    @Column(name = "processing_start_dt")
    private Timestamp processingStartDt;
    @Column(name = "processing_end_dt")
    private Timestamp processingEndDt;
    @Column(name = "attributes")
    private String attributes;
    @Column(name = "subject_id")
    private String subjectId;
    @Column(name = "exception_message")
    private String exceptionMessage;
    @Column(name = "message_type")
    private String messageType;
    @Column(name = "message_guid")
    private String messageGuid;
    @Column(name = "previous_message_guid")
    private String previousMessageGuid;
    @Column(name = "external_patient_id")
    private String externalPatientId;
    @Column(name = "portal_consumer_id")
    private Long portalConsumerId;
    @Column(name = "external_provider_id")
    private String externalProviderId;
    @Column(name = "portal_staff_id")
    private Long portalStaffId;
    @Column(name = "external_message_id")
    private String externalMessageId;
    @Column(name = "revision")
    private Integer revision;
    @Column(name = "error_cd")
    private String errorCd;
    @Column(name = "error_description")
    private String errorDescription;
    @Column(name = "error_severity")
    private String errorSeverity;

    @ManyToOne
    @JoinColumn(name = "processing_status_type_cd")
    private ProcessingStatusTypeEntity processingStatusTypeEntity;

    @OneToMany(mappedBy = "messageEntity")
    private List<ActivityEntity> activityEntityList;

    @OneToMany(mappedBy = "messageEntity")
    private List<ObjectRefEntity> objectRefEntityList;
}
