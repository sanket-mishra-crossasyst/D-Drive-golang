package com.crossasyst.activities.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;
import java.sql.Timestamp;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DataJob {
    private Long dataJobId;
    private String jobDirection;
    private String dataJobGuid;
    private String inputFileName;
    private Timestamp processingStartDt;
    private Timestamp processingEndDt;
    private String dataFeed;
    private String dataSource;
    private String dataPartner;
    private String msgType;
    private String jobType;
    private String externalSystemName;
    private Long orgId;
    private String orgUuid;
}
