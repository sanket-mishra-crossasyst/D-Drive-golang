package com.crossasyst.activities.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.sql.Timestamp;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "data_job")
public class DataJobEntity {

    @Id
    @Column(name = "data_job_id")
    private Long dataJobId;
    @Column(name = "job_direction")
    private String jobDirection;
    @Column(name = "data_job_guid")
    private String dataJobGuid;
    @Column(name = "input_file_name")
    private String inputFileName;
    @Column(name = "processing_start_dt")
    private Timestamp processingStartDt;
    @Column(name = "processing_end_dt")
    private Timestamp processingEndDt;
    @Column(name = "data_feed")
    private String dataFeed;
    @Column(name = "data_source")
    private String dataSource;
    @Column(name = "data_partner")
    private String dataPartner;
    @Column(name = "msg_type")
    private String msgType;
    @Column(name = "job_type")
    private String jobType;
    @Column(name = "external_system_name")
    private String externalSystemName;
    @Column(name = "org_id")
    private Long orgId;
    @Column(name = "org_uuid")
    private String orgUuid;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "data_channel")
    private DataChannelEntity dataChannelEntity;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "job_status_type_cd")
    private JobStatusTypeEntity jobStatusTypeEntity;


}
