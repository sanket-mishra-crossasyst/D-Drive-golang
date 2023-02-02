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
@Table(name = "activity")
public class ActivityEntity {

    @Id
    @Column(name = "activity_id")
    private Long activityId;
    @Column(name = "activity_name")
    private String activityName;
    @Column(name = "processing_start_date")
    private Timestamp processingStartDate;
    @Column(name = "processing_end_date")
    private Timestamp processingEndDate;
    @Column(name = "revision")
    private Long revision;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "msg_id")
    private MessageEntity messageEntity;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "activity_type")
    private ActivityTypeEntity activityTypeEntity;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "processing_status_type")
    private ProcessingStatusTypeEntity processingStatusTypeEntity;

}
