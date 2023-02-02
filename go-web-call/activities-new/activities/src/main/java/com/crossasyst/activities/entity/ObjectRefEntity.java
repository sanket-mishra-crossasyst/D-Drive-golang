package com.crossasyst.activities.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "object_ref")
public class ObjectRefEntity {

    @Id
    @Column(name = "object_ref_id")
    private Long objectRefId;
    @Column(name = "object_ref")
    private String objectRef;
    @Column(name = "msg_guid")
    private String msgGuid;
    @Column(name = "revision")
    private Integer revision;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "msg_id")
    private MessageEntity messageEntity;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "node_type_cd")
    private NodeTypeEntity nodeTypeEntity;


}
